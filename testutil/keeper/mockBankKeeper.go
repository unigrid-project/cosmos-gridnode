package keeper

// Import the necessary packages
import (
	"context"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/unigrid-project/cosmos-gridnode/x/gridnode/types"
)

// Define a struct that implements the types.BankKeeper interface
type MockBankKeeper struct {
	// You can include any additional fields or methods needed for your tests
}

// Implement the SendCoinsFromAccountToModule method
func (m *MockBankKeeper) SendCoinsFromAccountToModule(ctx context.Context, from sdk.AccAddress, to string, coins sdk.Coins) error {
	// Mock logic for SendCoinsFromAccountToModule
	return nil // Return nil to indicate success
}

// Implement the AddCoins method
func (m *MockBankKeeper) AddCoins(ctx context.Context, addr sdk.AccAddress, amt sdk.Coins) error {
	// Mock logic for AddCoins
	return nil // Return nil to indicate success
}

// Implement the missing SendCoinsFromModuleToAccount method
func (m *MockBankKeeper) SendCoinsFromModuleToAccount(ctx context.Context, module string, addr sdk.AccAddress, amt sdk.Coins) error {
	// Mock logic for SendCoinsFromModuleToAccount
	return nil // Return nil to indicate success
}

// Implement the missing GetBalance method
func (m *MockBankKeeper) GetBalance(ctx context.Context, addr sdk.AccAddress, denom string) sdk.Coin {
	// Mock logic for GetBalance
	return sdk.NewCoin("token", math.NewInt(100))
}

// Implement the SpendableCoins method
func (m *MockBankKeeper) SpendableCoins(ctx context.Context, addr sdk.AccAddress) sdk.Coins {
	// Mock logic for SpendableCoins, return an example of spendable coins
	return sdk.NewCoins(sdk.NewCoin("token", math.NewInt(100)))
}

// Create a function to create an instance of the mock BankKeeper
func NewMockBankKeeper() types.BankKeeper {
	return &MockBankKeeper{}
}
